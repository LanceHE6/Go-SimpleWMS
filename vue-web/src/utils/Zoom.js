class Zoom {
    // 监听方法兼容写法
    addHandler(element, type, handler) {
        if (element.addEventListener) {
            element.addEventListener(type, handler, false);
        } else if (element.attachEvent) {
            element.attachEvent('on' + type, handler);
        } else {
            element['on' + type] = handler;
        }
    }
    // 校正浏览器缩放比例
    correct() {
        // 当前页面屏幕分辨率
        let width = document.documentElement.clientWidth
        let height = document.documentElement.clientHeight
        //1490 为你的设计稿的宽度
        document.getElementsByTagName('body')[0].style.zoom = height * 2 / width ;

    }
    // 监听页面缩放
    watch() {
        this.addHandler(window, 'resize', function() { // 注意这个方法是解决全局有两个window.resize
            // 重新校正
            this.correct()
        })
    }
    // 初始化页面比例
    init() {
        // 初始化页面校正浏览器缩放比例
        this.correct();
        // 开启监听页面缩放
        this.watch();
    }
}
export default Zoom;