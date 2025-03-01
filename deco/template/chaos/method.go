package chaos

var Method = `func (i *{{implementation}}) {{method_signature}} {
    i.applyDelay()

    err := i.getErr()
    if err != nil {
        {{method_return}}
    }

    {{method_call}}
}
`
