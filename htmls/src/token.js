const getToken = ()=>{
    let token = localStorage.getItem('token');
    return token === null ? '': token;
}

const deleteToken = ()=>{
    localStorage.removeItem('token');
}

const setToken =(token)=>{
    localStorage.setItem('token',token);
}


export {getToken,deleteToken,setToken}