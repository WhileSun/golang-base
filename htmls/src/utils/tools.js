import requests from '@/request/index';
import {message} from 'antd';


function breakWords(str,liheHeight=12){
	return showHtml('<div style="line-height:'+liheHeight+'px;">'+str.replace(/\|/g,'<br>')+'</div>');
}

function showHtml(context ){
    let html = { __html:context };
    return (<div dangerouslySetInnerHTML={html}></div>);
}

function S4() {
    return (((1+Math.random())*0x10000)|0).toString(16).substring(1);
}

function getRandStr(prefix="") {
    var uuid = prefix+S4()+S4();
    return uuid;
}

function getData(api,apiParams,callback,succTips=false){
    return requests.post(api, apiParams).then(function (resp) {
        // console.log('resp',resp);
        if(resp.code ==0){
            if(succTips){
                console.log(succTips);
                message.success(resp.msg);
            }
            if(callback){
                return callback(resp.data);
            }
        }else{
          message.error(resp.msg);
        }
      }).catch(function(error){
        // message.error("api error");
        console.log('error',error);
      });
}

function inArray(val,arr,types){
    if(types == 'int'){
        val = parseInt(val);
    }
    return arr.includes(val);
}

function arrTransName(arr,keys){
    let newArrs = [];
    arr.map((val)=>{
        let newArr = {};
        for(let key in keys){
           newArr[keys[key]] = (val[key]).toString();
        }
        newArrs.push(newArr);
    });
    return newArrs;
}

function arrTransObj(arr,key,valKey){
    let newObjs = {}
    arr.map((val)=>{
        newObjs[val[key]] = val[valKey];
    });
    return newObjs;
}

function toTree(data){
     // 删除 所有 children,以防止多次调用
    data.forEach(function (item) {
        delete item.children;
    });

    // 将数据存储为 以 id 为 KEY 的 map 索引数据列
    var map = {};
    data.forEach(function (item) {
        map[item.id] = item;
    });
    //console.log(map);
    var val = [];
    data.forEach(function (item) {
        // 以当前遍历项，的pid,去map对象中找到索引的id
        var parent = map[item.parent_id];
        // 好绕啊，如果找到索引，那么说明此项不在顶级当中,那么需要把此项添加到，他对应的父级中
        if (parent) {
            (parent.children || ( parent.children = [] )).push(item);
        } else {
            //如果没有在map中找到对应的索引ID,那么直接把 当前的item添加到 val结果集中，作为顶级
            val.push(item);
        }
    });
    return val;
}

export {breakWords,showHtml,getRandStr,getData,inArray,arrTransName,arrTransObj,toTree}