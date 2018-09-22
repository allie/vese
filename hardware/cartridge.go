package hardware

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

var KnownGames = map[string]string{
	"54cf17c48707467295749490d458eafb": "Demonstration Cartridge",
	"f6916b665893aa8138cde57c37e50ada": "Demonstration Cartridge 2",
	"7e5c26a6d1f9a90c68669a9800ba522d": "Videocart-01: Tic-Tac-Toe, Shooting Gallery, Doodle, Quadra-Doodle",
	"4f11f13cbca685cb20e888f87b3b1586": "Videocart-02: Desert Fox, Shooting Gallery",
	"b074c867f235fb69ced96c6916673b45": "Videocart-03: Video Blackjack",
	"1b409fe1154584f4d1ab76b344a73d99": "Videocart-04: Spitfire",
	"32cca8ff09041a39251d7aade21ee22f": "Videocart-05: Space War",
	"a8e6103fcae4d0f9e14d9edcfc3fc493": "Videocart-06: Math Quiz I (Addition & Subtraction)",
	"86b77eafdf7b806e19e01724987e384f": "Videocart-07: Math Quiz II (Multiplication & Division)",
	"4c10fa5c7316c59efa241043fc67dfa8": "Videocart-08: Magic Numbers (Mind Reader & Nim)",
	"f80af74b09d058b90e719bb7dfbdd50e": "Videocart-09: Drag Race",
	"6565df74539476d66fd78de1bac0259c": "Videocart-10: Maze, Jailbreak, Blind man's bluff, Trailblazer",
	"53e4cc2da0f2c167e0692b794cb7692c": "Videocart-10: Maze, Jailbreak, Blind man's bluff, Trailblazer (Alt 1)",
	"d89b48ae8c906488caac2b2ae1d63d88": "Videocart-11: Backgammon, Acey-Deucey",
	"4fa83f734c139963aa02bdbb7a52e500": "Videocart-12: Baseball",
	"3783b6ac359e21b99cfa17773aa811c6": "Videocart-13: Robot War, Torpedo Alley",
	"4cb12edae37df23851884b82ca410754": "Videocart-14: Sonar Search",
	"2b3ca549e27579e4519a765fd8f52d0f": "Videocart-15: Memory Match 1, Memory Match 2",
	"6ffedaed3c5cd8ba74d98901849cc451": "Videocart-16: Dodge' It",
	"e90339b7068c6227d54f3c0ca637e017": "Videocart-17: Pinball Challenge",
	"5cbcda1c44f0dad02b0dfe209b6325d5": "Videocart-17: Pinball Challenge (Alt 1)",
	"0124cd0b61df5502aabd59029ccb6d5a": "Videocart-18: Hangman",
	"35d61d40ef7ec337cba092aabac74dbd": "Videocart-19: Checkers",
	"90a9b3952568f91502a7088bfb0ae07e": "Videocart-20: Video Whizball",
	"25e231e7a464a32b4715bfb47af89240": "Videocart-21: Bowling",
	"dfb66ee145fab65062fdeadafc8dc34c": "Videocart-22: Slot Machine",
	"9e0711b140e22729687db1e1354980ab": "Videocart-23: Galactic Space Wars",
	"9a894d745356a050f95410983c3bc54a": "Videocart-24: Pro Football",
	"bb7f7bbbe21f142591cdcafa98d7f6e4": "Videocart-25: Casino Poker",
	"f7bf7d55a7660ffa80d08ad1ba903ff7": "Videocart-26: Alien Invasion",
}

type Cartridge struct {
	Path     string
	Data     []byte
	Size     int
	Checksum string
	Title    string
	Known    bool
}

func NewCartridge(path string) (*Cartridge, error) {
	c := new(Cartridge)
	c.Path = path

	var err error
	c.Data, err = ioutil.ReadFile(c.Path)

	if err != nil {
		return nil, err
	}

	c.Size = len(c.Data)
	c.Checksum = fmt.Sprintf("%x", md5.Sum(c.Data))

	if title, ok := KnownGames[c.Checksum]; ok {
		c.Title = title
		c.Known = true
	} else {
		c.Title = filepath.Base(c.Path)
		c.Known = false
	}

	return c, nil
}
