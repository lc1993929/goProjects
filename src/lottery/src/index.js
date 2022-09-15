import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';
//直接引入会跟打包有冲突
// import 'semantic-ui-css/semantic.min.css';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
    <React.StrictMode>
        <App/>
    </React.StrictMode>
);

