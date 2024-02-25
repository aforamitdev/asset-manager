import { useState } from 'react';
import { Greet } from '../wailsjs/go/main/App';
import { Accordion } from '@radix-ui/react-accordion';
import { AccordionContent } from './shadcn/accordion';
import SideBar from './app/shared/sidebar/SideBar';

function App() {
  const [resultText, setResultText] = useState(
    'Please enter your name below 👇'
  );
  const [name, setName] = useState('');
  const updateName = (e: any) => setName(e.target.value);
  const updateResultText = (result: string) => setResultText(result);

  function greet() {
    Greet(name).then(updateResultText);
  }

  return (
    <div id='App' className='bg-white h-screen w-full'>
      <SideBar />
    </div>
  );
}

export default App;
