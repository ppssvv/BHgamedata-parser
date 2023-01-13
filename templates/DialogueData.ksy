meta:
  id: dialogue_data
  endian: le
seq:
  - id: file_size
    type: u4
  - id: num_entries
    type: s4
  # - id: junk
  #   type: u4
  #   repeat: expr
  #   repeat-expr: num_entry
  # - id: junk2
  #   type: u4
  #   repeat: expr
  #   repeat-expr: num_entry
  # - id: entry
  #   type: entry
  #   repeat: expr
  #   repeat-expr: num_entry
instances:
  entries:
    pos: 4 + 4 + (4 * _root.num_entries) + (4 * _root.num_entries)
    type: entry
    repeat: expr
    repeat-expr: _root.num_entries
types:
  u4_arr:
    seq:
      - id: num_entries
        type: s4
      - id: entries
        type: u4
        repeat: expr
        repeat-expr: num_entries
  str_arr:
    seq:
      - id: num_entries
        type: s4
      - id: entries
        type: str_with_len
        repeat: expr
        repeat-expr: num_entries
  str_with_len:
    seq:
      - id: len
        type: u2
      - id: value
        type: str
        encoding: UTF-8
        size: len
  dialogue_entry:
    seq:
      - id: len
        type: s2
      - id: value
        type: str
        encoding: UTF-8
        size: len
      - id: duration
        type: f4
  dialogue_content:
    seq:
      - id: num_dialogues
        type: s4
      - id: dialogues
        type: dialogue_entry
        repeat: expr
        repeat-expr: num_dialogues
  entry_header:
    seq:
      - id: entry_id
        type: s4
      - id: addrto_postdialogue
        type: u4
      - id: unk1
        type: s4
      - id: addrto_predialogue
        type: u4
      - id: dialogue_type
        type: s1
      - id: unk2
        type: s4
      - id: unk3
        type: s4
      - id: addrto_avatar
        type: u4
      - id: avatar_id
        type: s4
      - id: dress_id
        type: s4
      - id: avatar_vice_key
        type: s4
      - id: screen_side
        type: s1
      - id: face_id
        type: s1
      - id: addrto_face
        type: u4
      - id: addrto_face_animation
        type: u4
      - id: addrto_face_effect
        type: u4
      - id: animation_id
        type: s1
      - id: distortion
        type: s1
      - id: transparency
        type: f4
      - id: addrto_smth2
        type: u4
      - id: addrto_smth3
        type: u4
      - id: addrto_smth4
        type: u4
      - id: unk5
        type: s4
      - id: addrto_content
        type: u4
      - id: addrto_smth5
        type: u4
      - id: addrto_smth6
        type: u4
      - id: background_effect
        type: s1
      - id: enter_effect
        type: s1
      - id: addrto_audio_id
        type: u4
      - id: audio_delay
        type: f4
      - id: addrto_lip_motion
        type: u4
      - id: unk7
        type: s1
      - id: addrto_smth7
        type: u4
      - id: addrto_smth8
        type: u4
      - id: addrto_smth9
        type: u4
      - id: addrto_smth10
        type: u4
      - id: unk8
        type: s4
      - id: unk9
        type: s4
      - id: addrto_smth11
        type: u4
      - id: addrto_smth12
        type: u4
      - id: addrto_smth13
        type: u4
  entry_data:
    seq: 
      - id: post_dialogue
        type: u4_arr
      - id: pre_dialogue
        type: u4_arr
      - id: avatar
        type: str_with_len
      - id: face
        type: str_with_len
      - id: face_animation
        type: str_with_len
      - id: face_effect
        type: str_with_len
      - id: unknown_6
        type: str_with_len
      - id: unknown_7
        type: str_arr
      - id: unknown_8
        type: str_with_len
      - id: content
        type: dialogue_content
      - id: unknown_9
        type: str_with_len
      - id: unknown_10
        type: str_with_len
      - id: audio_id
        type: str_with_len
      - id: lip_motion
        type: str_with_len
      - id: unknown_11
        type: str_with_len
      - id: unknown_12
        type: str_with_len
      - id: unknown_13
        type: str_arr
      - id: unknown_14
        type: str_arr
      - id: unknown_15
        type: u4_arr
      - id: unknown_16
        type: str_with_len
      - id: unknown_17
        type: str_with_len
  entry:
    seq:
      - id: header
        type: entry_header
      - id: data
        type: entry_data
