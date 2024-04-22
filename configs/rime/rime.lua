
-- processors:

-- 以词定字，可在 default.yaml → key_binder 下配置快捷键，默认为左右中括号 [ ]
select_character = require("select_character")

-- Unicode，U 开头
unicode = require("unicode")

-- 数字、人民币大写，R 开头
number_translator = require("number_translator")

-- translators:

-- 农历，可在方案中配置触发关键字。
lunar = require("lunar")

-- 日期时间，可在方案中配置触发关键字。
date_translator = require("date_translator")

-- filters:

-- 错音错字提示
-- 关闭此 Lua 时，同时需要关闭 translator/spelling_hints，否则 comment 里都是拼音
corrector = require("corrector")

-- 暴力 GC
-- 详情 https://github.com/hchunhui/librime-lua/issues/307
function force_gc()
    -- collectgarbage()
    collectgarbage("step")
end