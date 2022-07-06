import $ from 'jquery';
//常规fiexd
export const normalResize = (divId, footHeight) => {
    if (checkTableShow(divId)) {
        let elem = $('.' + divId);
        let thtop = elem.find('.ant-table-body:first').offset().top;
        elem.find('.ant-table-body').height($(window).height() - thtop - footHeight);
    } else {
        setTimeout(() => {
            normalResize(divId, footHeight)
        }, 100);
    }
}

//判断是否加载
export const checkTableShow = (divId) => {
    let elem = $('.' + divId);
    if (elem.find('.ant-table-body').offset() === undefined) {
        return false;
    } else {
        return true;
    }
}

//模态框的fiexd
export const modalResize = (divId, footHeight) => {
  let elem = $('.' + divId);
  let thtop = elem.find('.ws-table-header').height();
  let thead = elem.find('.ant-table-header thead').height();
  elem.find('.ant-table-body').height($(window).height() - 180 - thtop - thead - 20 - footHeight);
}