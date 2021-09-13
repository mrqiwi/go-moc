package mocp

import (
	"errors"
	"fmt"
	"os/exec"
)

var (
	errCannotRunCommand = errors.New("cannot run command")
	errUnknownCommand   = errors.New("unknown command")
)

type MOCp struct{}

func NewMOCp() *MOCp {
	return &MOCp{}
}

func (m MOCp) Play() error {
	cmd := exec.Command("mocp", "--play")

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%w: %s", errCannotRunCommand, err)
	}

	return nil
}

func (m MOCp) Stop() error {
	cmd := exec.Command("mocp", "--stop")

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%w: %s", errCannotRunCommand, err)
	}

	return nil
}

func (m MOCp) TogglePause() error {
	cmd := exec.Command("mocp", "--toggle-pause")

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%w: %s", errCannotRunCommand, err)
	}

	return nil
}

func (m MOCp) Next() error {
	cmd := exec.Command("mocp", "--next")

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%w: %s", errCannotRunCommand, err)
	}

	return nil
}

func (m MOCp) Seek(isForward bool) error {
	var cmd *exec.Cmd

	if isForward {
		cmd = exec.Command("mocp", "--seek", "5")
	} else{
		cmd = exec.Command("mocp", "--seek", "-5")
	}

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%w: %s", errCannotRunCommand, err)
	}

	return nil
}

func (m MOCp) Previous() error {
	cmd := exec.Command("mocp", "--previous")

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%w: %s", errCannotRunCommand, err)
	}

	return nil
}

func (m MOCp) Exit() error {
	cmd := exec.Command("mocp", "--exit")

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%w: %s", errCannotRunCommand, err)
	}

	return nil
}

func (m MOCp) Volume(isIncrease bool) error {
	var cmd *exec.Cmd

	if isIncrease {
		cmd = exec.Command("mocp", "--volume", "+5")
	} else{
		cmd = exec.Command("mocp", "--volume", "-5")
	}

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%w: %s", errCannotRunCommand, err)
	}

	return nil
}

func (m MOCp) Mute() error {
	cmd := exec.Command("mocp", "--volume", "0")

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%w: %s", errCannotRunCommand, err)
	}

	return nil
}

func (m MOCp) TrackInfo() (TrackInfo, error) {
	var trackInfo = TrackInfo{}

	cmd := exec.Command("mocp", "--info")

	out, err := cmd.Output()
	if err != nil {
		return trackInfo, fmt.Errorf("%w: %s", errCannotRunCommand, err)
	}

	return fillTrackInfo(string(out)), nil
}
