/**
 * 初始化table 字段展示筛选框
 * @param {Array} fields 
 * @returns 
 */
const initShowTreeFunc = (fields) => {
  let allData = [];
  let allKeys = [];
  for (var i = 0, len = fields.length; i < len; i++) {
    let d = fields[i];
    let title = d['title'];
    let newtitle = '';
    if (title instanceof Object) {
      newtitle =
        title['props']['dangerouslySetInnerHTML']['__html'].match(
          /<div.*?>(.*?)<\/div>/,
        );
      title = newtitle[1].replace('<br>', '');
    }
    allData.push({ title: title, key: d['name'] + '_' + i });
    allKeys.push(d['name'] + '_' + i);
  }
  return { allData: allData, allKeys: allKeys };
};
export default initShowTreeFunc;