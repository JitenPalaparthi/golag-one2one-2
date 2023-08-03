
// calculator

func Calc(a,b any)(sum,sub,mul,div,mod any,err error){

}

// it should work only for numbers

Calc(10,20)

30,-10,200,.5,0,nil

Calc(false,20)
0,0,0,0,0,"error in input"

Calc("Hello","World")
0,0,0,0,0,"error in input"