import { useRef } from 'react';
const useForm = (form) => {
  const formRef = useRef();
  if (!formRef.current) {
    if (typeof (form) == 'object' ) {
      formRef.current = form;
    } else {
      formRef.current = {};
    }
  }
  return formRef.current;
};

export default useForm;
