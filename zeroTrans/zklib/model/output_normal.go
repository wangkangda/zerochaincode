package model

type NormalOutput struct{
    receiver    Address
    amount      int
}

func (o *NormalOutput)GetType(){
    return NormalTransaction
}

func (o *NormalOutput)Prepare(ctx Context){
    ctx.AddAmount( o.receiver )
}

func (o *NormalOutput)Verify(context Context )bool{
    return true
}

func (o *NormalOutput)Execute(context Context)error{
    context.amount[ o.receiver ] += amount
}

func (o *NormalOutput)String()string{
    return fmt.Sprintf("%v\n%v", o.receiver.String(), o.amount)
}

func (o *NormalOutput)FromString(istr string)error{
    ostr := SplitN(istr, "\n")
    err := o.receiver.FromString(ostr[0])
    if err != nil{
        return err
    }
    o.amount, err = strconv.Atoi(ostr[1])
    return err
}
