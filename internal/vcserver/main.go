package vcserver

import "github.com/Pauloo27/vcserver/internal/models"

type VoiceServer struct {
	Target models.AbstractTarget
}

func NewVoiceServer(target models.AbstractTarget) *VoiceServer {
	return &VoiceServer{
		Target: target,
	}
}

func (s *VoiceServer) Start() error {
	return s.Target.Start()
}

func (s *VoiceServer) Stop() error {
	return s.Target.Stop()
}
