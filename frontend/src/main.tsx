import React from 'react';
import { createRoot } from 'react-dom/client';
import './output.css';
import App from './App';
import '@radix-ui/themes/styles.css';
import { Theme } from '@radix-ui/themes';

const container = document.getElementById('root');

const root = createRoot(container!);

root.render(
  <React.StrictMode>
    <Theme
      accentColor='crimson'
      // appearance='dark'
      grayColor='sand'
      radius='large'
      scaling='95%'
    >
      <App />
    </Theme>
  </React.StrictMode>
);
