require("options")
require("keymaps")
require("plugins")
require('colorscheme')
if next(vim.fn.argv()) ~= nil then
    require('lualine').setup({})
end
