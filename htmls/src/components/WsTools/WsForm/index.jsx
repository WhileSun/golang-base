import { forwardRef } from 'react';
import WsForm from './Form';
import useForm from './hooks/useForm';

var InternalForm = forwardRef(WsForm);
var RefWsForm = InternalForm;
RefWsForm.useForm = useForm;
export { useForm };
export default RefWsForm;
