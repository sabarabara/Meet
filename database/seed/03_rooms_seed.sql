INSERT INTO rooms (
  roomid, recruitid, userid, role, isfinish
) VALUES
(
  '746c5625-d9c1-4fad-bbe5-f8026a0482a1',
  'e2f46b58-a5d2-4376-a2f3-3df243b2c2e5',
  'fcf49967-0058-4051-a704-22bd99078606',
  'owner',
  FALSE
)ON CONFLICT (roomid) DO NOTHING;