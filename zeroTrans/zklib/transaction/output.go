package transaction

type Output struct{
}

func (o *Output) Verify () error{
    return nil
}

func (o *Output) Execute(){
}

