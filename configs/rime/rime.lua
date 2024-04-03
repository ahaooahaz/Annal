-- 两个短横杆代表 lua 的注释行

-- 翻译器：自动转换日期时间
function date_translator(input, seg)
   if (input == "date") or (input == "time") then
      --- Candidate(type, start, end, text, comment)
      yield(Candidate("date", seg.start, seg._end, os.date("%Y年%m月%d日"), "日期"))
      yield(Candidate("time", seg.start, seg._end, os.date("%H:%M:%S"), "时间"))
   end
end