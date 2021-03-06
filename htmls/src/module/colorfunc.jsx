import colors from '@/module/colors';
import {Tag } from 'antd';

//状态
export const statusFunc = (v,statusList)=>{
    const typeColors = {'true':colors.blue_6,'false':colors.volcano_6};
    return <Tag color={typeColors[v]} key={v}>{statusList[v]}</Tag>
  }

export const menuTypeFunc = (v,MenuTypeList)=>{
    const typeColors = {1:colors.geekblue_6,2:colors.orange_6};
    return <Tag color={typeColors[v]}>{MenuTypeList[v]}</Tag>
}

export const workTaskLevelFunc=(row)=>{
    const typeColors = {1:colors.green_6,2:colors.orange_6,3:colors.red_6};
    for(let key in row){
        row[key] = <Tag color={typeColors[key]}>{row[key]}</Tag> 
    }
    return row;
}

export const workTaskStatusFunc=(key,row)=>{
    const typeColors = {1:colors.grey_6,2:colors.blue_6};
    if(key>2){
        return <Tag color={colors.gold_6}>{row[key]}</Tag>
    }else{
        return <Tag color={typeColors[key]}>{row[key]}</Tag>
    }
}