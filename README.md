# Bignum - arbitrary precision decimal arithmetic for lua

This package extends [gopher-lua](https://github.com/yuin/gopher-lua) with arbitrary precision decimal arithmetic

```lua
local bignum = require("bignum").new

local fv = 0.1578416574
local loops = 3
local smx = 0
local sm =  bignum(0)
for i=1,loops,1 do
	sm = sm:add(bignum(fv))
	smx = smx + fv
end
print(smx:string())
```

