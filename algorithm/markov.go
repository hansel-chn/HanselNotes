package algorithm

import (
	"bufio"
	"bytes"
	"io"
	"math"
)

type Classifier struct {
	counts  map[byte]map[byte]float64
	allowed map[byte]bool
}

func (c *Classifier) Train(r io.Reader) error {
	scanner := bufio.NewScanner(r)

	// count the transaction from a->b
	for scanner.Scan() {
		in := bytes.TrimSpace(scanner.Bytes())
		for i := 0; i < len(in)-1; i++ {
			a, b := in[i], in[i+1]
			if !c.allowed[a] || !c.allowed[b] {
				continue
			}
			c.counts[a][b]++
		}
	}

	// normalize to log probabilities.
	for a, transitions := range c.counts {
		sum := 0.0
		for _, counts := range transitions {
			sum += counts
		}
		for b, counts := range transitions {
			c.counts[a][b] = math.Log(counts / sum)
		}
	}

	return nil
}

func (c *Classifier) Analyze(junk []byte) float64 {
	return c.avg(junk)
}

// avg is the average transition probability for a slice of bytes.
func (c *Classifier) avg(in []byte) float64 {
	in = bytes.TrimSpace(in)
	count := 0
	log := 0.0
	i := 0
	for ; i < len(in)-1; i++ {
		if c.allowed[in[i]] {
			break
		}
	}
	a := in[i]
	for j := i + 1; j < len(in); j++ {
		b := in[j]
		if !c.allowed[b] {
			continue
		}
		log += c.counts[a][b]
		a = b
		count++
	}
	if count == 0 {
		return 0
	}
	return math.Exp(log / float64(count))
}

func NewClassifier(allowedChars []byte) *Classifier {
	classifier := &Classifier{allowed: map[byte]bool{}, counts: map[byte]map[byte]float64{}}
	for _, c := range allowedChars {
		classifier.allowed[c] = true
	}
	for r := range classifier.allowed {
		classifier.counts[r] = map[byte]float64{}
		for k := range classifier.allowed {
			// init value, assume these chars have appeared a few times
			classifier.counts[r][k] = 10
		}
	}

	return classifier
}
