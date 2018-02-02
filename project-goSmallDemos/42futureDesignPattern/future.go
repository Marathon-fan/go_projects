package future

type SuccessFunc func(string)
type Failfunc func(error)
type ExecuteStringFunc func(string, error)

type MaybeString struct {

}

type (s *MaybeString) Success(f SuccessFunc) MaybeString{
	return nil	
}

type (s *MaybeString) Fail(f FailFunc) MaybeString{
	return nil	
}

type (s *MaybeString) Execute(f ExecuteStringFunc) {
}