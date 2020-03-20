package server

import (
	"go-lib/log"
	"go-lib/utils"
	"strings"
	"voip/rw"
)

func (s *Server) Auth(readerWriter rw.ReaderWriterCloser) (uid int64, rid int32, ok bool) {
	var header = make([]byte, 2+32+4) //(token + roomid)
	_, err := readerWriter.Read(header)
	if err != nil {
		log.Error(err)
		return
	}

	var token = strings.TrimSpace(string(header[2 : 32+2]))
	rid = utils.BytesToInt32(header[32+2:])

	uid, ok = s.GetToken(token)
	if !ok { //鉴权
		log.Warnf("access denied, uid:%d", uid)
		return
	}
	r, ok := s.GetRoom(rid)
	if !ok {
		log.Warnf("access denied, uid:%d", uid)
		return
	}

	if !r.JoinRoom(uid, readerWriter) {
		log.Warnf("access denied, roomId:%d, uid:%d", rid, uid)
		return
	}

	//成功连接 发送给chatserver
	s.JoinRoom(uid, rid, readerWriter.Name())
	return uid, rid, true
}
