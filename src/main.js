import Vue from 'vue'
import App from './App.vue'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';

Vue.use(ElementUI);
// Vue.config.delimiters = ["${","}"];
// Vue.config.delimiters = ["${","}"];
// Vue.config.unsafeDelimiters = ["${","}"];
Vue.options.delimiters = ['${', '}'];

new Vue({
    // delimiters: ['${', '}$'],
    el: '#app',
    delimiters: ['${', '}'],
    render: h => h(App)
});

// new Vue({
//     delimiters: ['${', '}'],
//     el: '#app',
//     router,
//     data:{
//         msg:'test'
//     },
//     template: '<div>${msg}</div>',
//     components: { App }
// })