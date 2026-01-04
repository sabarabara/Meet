INSERT INTO messages (messageid, roomid, senderid, content, isread) VALUES
  ('11111111-1111-1111-1111-111111111111', '746c5625-d9c1-4fad-bbe5-f8026a0482a1', 'fcf49967-0058-4051-a704-22bd99078606', 'Hello, this is the first message', FALSE)
ON CONFLICT (messageid) DO NOTHING;
