// Code generated by "stringer -type=player"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[dealer-0]
	_ = x[player1-1]
	_ = x[player2-2]
}

const _player_name = "dealerplayer1player2"

var _player_index = [...]uint8{0, 6, 13, 20}

func (i player) String() string {
	if i < 0 || i >= player(len(_player_index)-1) {
		return "player(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _player_name[_player_index[i]:_player_index[i+1]]
}