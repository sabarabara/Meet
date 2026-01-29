INSERT INTO recruit (
  recruitid, userid, area, imgurl,
  man, woman, vacant_man, vacant_woman,
  comment, date
) VALUES (
  'e2f46b58-a5d2-4376-a2f3-3df243b2c2e5',
  'fcf49967-0058-4051-a704-22bd99078606',
  'Tokyo',
  'https://example.com/image.png',
  2, 1, 1, 0,
  'test comment',
  '2026-01-02'
)
ON CONFLICT (recruitid) DO NOTHING;
