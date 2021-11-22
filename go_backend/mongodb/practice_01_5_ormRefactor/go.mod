module practice_01_5_ormRefactor

go 1.16

require github.com/julienschmidt/httprouter v1.3.0

require "models" v0.0.0
replace "models" => "./models"
