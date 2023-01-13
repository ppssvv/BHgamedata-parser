meta:
  id: text_map
  endian: le
seq:
  - id: size
    type: s4
  - id: num_textlist
    type: s4
  # - id: hashlist
  #   type: hashentry
  #   repeat: expr
  #   repeat-expr: num_textlist
  # - id: textaddrlist
  #   type: u4
  #   repeat: expr
  #   repeat-expr: num_textlist
  # - id: textlist
  #   type: textentry
  #   repeat: expr
  #   repeat-expr: num_textlist
instances:
  textlist:
    pos: 4 + 4 + 5 * _root.num_textlist + 4 * _root.num_textlist
    type: textentry
    repeat: expr
    repeat-expr: num_textlist
types:
  # hashentry:
  #   seq:
  #     - id: key
  #       type: s1
  #     - id: hashaddr
  #       type: u4
  textentry:
    seq:
      - id: addr1
        type: u4
      - id: addr2
        type: u4
      - id: key
        type: u1
      - id: hash
        type: u4
      - id: length
        type: s2
      - id: text
        type: str
        size: length
        encoding: UTF-8
