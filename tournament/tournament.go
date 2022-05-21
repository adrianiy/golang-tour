// Package tournament contains methods to handle your tournament.
package tournament

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Team struct {
	Team           string
	MP, W, D, L, P int8
}

type League map[string]*Team

type Classification []Team

// win update data with win case
func (r *Team) win() {
	r.W += 1
	r.P += 3
}

// draw update team data with draw case
func (r *Team) draw() {
	r.D += 1
	r.P += 1
}

// loss update team data with loss case
func (r *Team) loss() {
	r.L += 1
}

// getTeam recovers team from league map.
func (l League) getTeam(name string) *Team {
	if l[name] == nil {
		l[name] = &Team{Team: name}
	}
	return l[name]
}

// updateResults udpate team data in league map.
func (l *League) updateResults(results []string) error {
	if len(results) < 3 {
		return errors.New("Invalid line.")
	}
	home := l.getTeam(results[0])
	visitor := l.getTeam(results[1])

	home.MP += 1
	visitor.MP += 1

	switch results[2] {
	case "win":
		home.win()
		visitor.loss()
	case "draw":
		home.draw()
		visitor.draw()
	case "loss":
		home.loss()
		visitor.win()
	default:
		return errors.New("Invalid result.")
	}

	return nil
}

// sort order classification table by points or name in case of draw
func (c Classification) sort() {
	sort.Slice(c, func(i, j int) bool {
		if c[i].P == c[j].P {
			return c[i].Team < c[j].Team
		}
		return c[i].P > c[j].P
	})
}

/*
* Tally method creates a classification table following example.
* MP: Matches Played
* W: Matches Won (3 points)
* D: Mathces Drawn (tied) (1 point)
* L: Matches Lost (0 points)
* P: Points
 */
func Tally(reader io.Reader, writer io.Writer) error {
	var buffer bytes.Buffer
	var buf = new(strings.Builder)
	var league = make(League)
	var classification Classification

	_, err := io.Copy(buf, reader)

	if err != nil {
		return err
	}

	lines := strings.Split(strings.ReplaceAll(strings.TrimSpace(buf.String()), "\r\n", "\n"), "\n")

	for _, line := range lines {
		if line != "" && line[0] != '#' {
			results := strings.Split(line, ";")
			err := league.updateResults(results)

			if err != nil {
				return err
			}
		}
	}

	for _, team := range league {
		classification = append(classification, *team)
	}

	classification.sort()

	buffer.WriteString(fmt.Sprintf("%-30s | %2s | %2s | %2s | %2s | %2s\n", "Team", "MP", "W", "D", "L", "P"))
	for _, cLine := range classification {
		buffer.WriteString(fmt.Sprintf("%-30s | %2d | %2d | %2d | %2d | %2d\n", cLine.Team, cLine.MP, cLine.W, cLine.D, cLine.L, cLine.P))
	}

	writer.Write(buffer.Bytes())

	return nil
}
