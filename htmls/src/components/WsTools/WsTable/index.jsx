import { forwardRef } from 'react';
import WsTable from './Table';
import useTable from './hooks/useTable';

var InternalForm = forwardRef(WsTable);
var RefWsTable = InternalForm;
RefWsTable.useTable = useTable;
export default RefWsTable;
