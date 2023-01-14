meta:
  id: dorm_event
  endian: le
seq:
  - id: size
    type: s4
  - id: num_entries
    type: s4
  # - id: junk1
  #   type: s4
  #   repeat: expr
  #   repeat-expr: num_entries
  # - id: junk2
  #   type: s4
  #   repeat: expr
  #   repeat-expr: num_entries
  # - id: entries
  #   type: entry
  #   repeat: expr
  #   repeat-expr: num_entries
instances:
  entries:
    pos: 4 + 4 + (4 * _root.num_entries) + (4 * _root.num_entries)
    type: entry
    repeat: expr
    repeat-expr: _root.num_entries
types:
  str_with_len:
    seq:
      - id: len
        type: u2
      - id: value
        type: str
        encoding: UTF-8
        size: len
  entry_header:
    seq:
      - id: hash
        type: s4
      - id: addr1
        type: u4
      - id: unk1
        type: s1
      - id: unk2
        type: f4
      - id: unk3
        type: f4
      - id: addr2
        type: u4
      - id: addr3
        type: u4
      - id: unk4
        type: s1
      - id: unk5
        type: s1
      - id: addr4
        type: u4
      - id: addr5
        type: u4
      - id: addr6
        type: u4
  entry_data:
    seq:
      - id: event_type
        type: str_with_len
      - id: emotion
        type: str_with_len
      - id: unk1
        type: str_with_len
      - id: unk2
        type: str_with_len
      - id: unk3
        type: str_with_len
      - id: unk4
        type: str_with_len
  entry:
    seq:
      - id: header
        type: entry_header
      - id: data
        type: entry_data