import React from 'react';
import { Button, Flex, Tabs } from '@radix-ui/themes';
type Props = {};

const OverViewSwitch = (props: Props) => {
  return (
    <Flex direction='column' gap='4'>
      <Tabs.Root defaultValue='world' className='py-3.5'>
        <Tabs.List size='2'>
          <Tabs.Trigger value='world'>World</Tabs.Trigger>
          <Tabs.Trigger value='IN'>INDIA</Tabs.Trigger>
          <Tabs.Trigger value='US'>US</Tabs.Trigger>
          <Tabs.Trigger value='CHINA'>CHINA</Tabs.Trigger>
          <Tabs.Trigger value='SGX'>SGX</Tabs.Trigger>
        </Tabs.List>

        <Tabs.Content value='world'>World</Tabs.Content>
        <Tabs.Content value='IN'>India</Tabs.Content>
        <Tabs.Content value='US'>US</Tabs.Content>
        <Tabs.Content value='CHINA'>CHINA</Tabs.Content>
        <Tabs.Content value='SGX'>CHINA</Tabs.Content>
      </Tabs.Root>
    </Flex>
  );
};

export default OverViewSwitch;
