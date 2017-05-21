package model

type Output interface{
    GetType() int
    Prepare(Context)
    Verify(Context)bool
    Execute(Context)error
    String()string
    FromString(string)error
}
