package vcserver

import "github.com/Pauloo27/vcserver/internal/models"

type VoiceServer struct {
	APIServers []models.AbstractAPIServer
	Target     models.AbstractTarget
}

func NewVoiceServer(target models.AbstractTarget, apiServers []models.AbstractAPIServer) *VoiceServer {
	return &VoiceServer{
		Target:     target,
		APIServers: apiServers,
	}
}

func (s *VoiceServer) Start() error {
	var err error
	if err = s.Target.Start(); err != nil {
		return err
	}
	for _, server := range s.APIServers {
		if err := server.Start(); err != nil {
			return err
		}
	}
	return nil
}

func (s *VoiceServer) Stop() error {
	return s.Target.Stop()
}
