package model

import (
	"errors"
	"math"
	"sync"

	"github.com/HackYourCareer/SmartKickers/internal/config"
	"github.com/HackYourCareer/SmartKickers/internal/controller/adapter"
)

type Game interface {
	AddGoal(int) error
	ResetScore()
	GetScore() GameScore
	GetScoreChannel() chan GameScore
	SubGoal(int) error
	IsFastestShot(float64) bool
	SaveFastestShot(adapter.ShotMessage)
	GetFastestShot() [config.HeatmapDimension][config.HeatmapDimension]int
	WriteToHeatmap(float64, float64) error
}

type game struct {
	score        GameScore
	scoreChannel chan GameScore
	m            sync.RWMutex
	fastestShot  adapter.ShotMessage
	heatmap      [config.HeatmapDimension][config.HeatmapDimension]int
}

type GameScore struct {
	BlueScore  int `json:"blueScore"`
	WhiteScore int `json:"whiteScore"`
}

func NewGame() Game {
	return &game{scoreChannel: make(chan GameScore, 32)}
}

func (g *game) ResetScore() {
	g.m.Lock()
	defer g.m.Unlock()
	g.score.BlueScore = 0
	g.score.WhiteScore = 0
	g.scoreChannel <- g.score
}

func (g *game) AddGoal(teamID int) error {
	g.m.Lock()
	defer g.m.Unlock()

	switch teamID {
	case config.TeamWhite:
		g.score.WhiteScore++
	case config.TeamBlue:
		g.score.BlueScore++
	default:
		return errors.New("bad team ID")
	}
	g.scoreChannel <- g.score

	return nil
}

func (g *game) GetScore() GameScore {
	g.m.RLock()
	defer g.m.RUnlock()

	return g.score
}

func (g *game) GetScoreChannel() chan GameScore {
	return g.scoreChannel
}

func (g *game) SubGoal(teamID int) error {
	g.m.Lock()
	defer g.m.Unlock()

	switch teamID {
	case config.TeamWhite:
		if g.score.WhiteScore > 0 {
			g.score.WhiteScore--
		}
	case config.TeamBlue:
		if g.score.BlueScore > 0 {
			g.score.BlueScore--
		}
	default:
		return errors.New("bad team ID")
	}
	g.scoreChannel <- g.score

	return nil
}

func (g *game) IsFastestShot(speed float64) bool {
	g.m.RLock()
	defer g.m.RUnlock()

	return g.fastestShot.Speed < speed
}

func (g *game) SaveFastestShot(msg adapter.ShotMessage) {
	g.m.Lock()
	defer g.m.Unlock()
	g.fastestShot.Speed = msg.Speed
	g.fastestShot.Team = msg.Team
}

func (g *game) GetFastestShot() [config.HeatmapDimension][config.HeatmapDimension]int {
	g.m.RLock()
	defer g.m.RUnlock()

	return g.heatmap
}

func (g *game) WriteToHeatmap(xCord float64, yCord float64) error {
	g.m.Lock()
	defer g.m.Unlock()

	x := int(math.Round(config.HeatmapDimension * xCord))
	y := int(math.Round(config.HeatmapDimension * yCord))
	if x > config.HeatmapDimension-1 || x < 0 {
		return errors.New("x cord out of index")
	}
	if y > config.HeatmapDimension-1 || y < 0 {
		return errors.New("y cord out of index")
	}
	g.heatmap[x][y]++
	return nil
}
