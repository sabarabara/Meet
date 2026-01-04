INSERT INTO users (
  userid, username, stars
) VALUES
  ('fcf49967-0058-4051-a704-22bd99078606', 'test_user', 4),
  ('c63fcb27-80ff-4a40-ae16-53465690df64', 'room_owner', 3),
  ('92c1e772-deca-4ab4-a075-60d8215553f5', 'room_member', 4),
  ('0a375ec0-d459-42f2-8160-f805baa0c163', 'review_writer', 5),
  ('2b34f2b5-799c-4046-83b8-96f385a5ac24', 'review_target', 4)
ON CONFLICT (userid) DO NOTHING;
