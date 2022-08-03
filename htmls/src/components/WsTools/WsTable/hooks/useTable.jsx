import { useRef } from 'react';

const useTable = (table) => { 
  const tableRef = useRef()
  if (!tableRef.current) {
    if (typeof (table) == 'object' ) {
      tableRef.current = table;
    } else {
      tableRef.current = {};
    }
  }
  return tableRef.current;
};

export default useTable;
