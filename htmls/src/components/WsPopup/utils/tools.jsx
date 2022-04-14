//参数是否存在
const paramIsset = (param, defaultValue = '') => {
  return param === undefined ? defaultValue : param;
};

export {
  paramIsset
};
