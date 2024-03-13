import OverViewSwitch from '@/app/market/header/OverViewSwitch';
import { cn } from '@/lib/utils';
import { AvatarFallback, AvatarImage, Avatar } from '@/shadcn/avatar';
import { Nav } from '@/shadcn/nav';
import {
  ResizableHandle,
  ResizablePanel,
  ResizablePanelGroup,
} from '@/shadcn/resizable';
import { Separator } from '@/shadcn/separator';
import { TooltipProvider } from '@/shadcn/tooltip';

import {
  ActivityIcon,
  Archive,
  Building,
  ReceiptIndianRupee,
  Trash2,
  Wallet,
  Command,
} from 'lucide-react';
import React from 'react';

type Props = {
  defaultLayout?: number[];
  navCollapsedSize?: number;
};

const SideBar = ({
  defaultLayout = [265, 440, 655],
  navCollapsedSize = 5,
}: Props) => {
  const [isCollapsed, setIsCollapsed] = React.useState(false);

  return (
    <TooltipProvider delayDuration={0}>
      <ResizablePanelGroup
        direction='horizontal'
        onLayout={(sizes: number[]) => {
          document.cookie = `react-resizable-panels:layout=${JSON.stringify(
            sizes
          )}`;
        }}
        className='h-full max-h-screen items-stretch'
      >
        <ResizablePanel
          defaultSize={defaultLayout[0]}
          collapsedSize={navCollapsedSize}
          collapsible={true}
          minSize={15}
          maxSize={20}
          onExpand={() => setIsCollapsed(false)}
          className={cn(
            isCollapsed &&
              'min-w-[50px] transition-all duration-300 ease-in-out'
          )}
          onCollapse={() => {
            setIsCollapsed(true);
            document.cookie = `react-resizable-panels:collapsed=${JSON.stringify(
              true
            )}`;
          }}
        >
          <div
            className={cn(
              'flex h-[51px] items-center',
              isCollapsed ? 'h-[51px] justify-center' : 'px-2'
            )}
          >
            <Avatar className='rounded-md '>
              <AvatarImage src='' className=' bg-gray-900 '></AvatarImage>
              <AvatarFallback className='bg-gray-500 '>AM</AvatarFallback>
            </Avatar>
          </div>
          <Separator />
          <Nav
            isCollapsed={isCollapsed}
            links={[
              {
                title: 'Overview',
                label: '',
                icon: Command,
                variant: 'default',
              },
              {
                title: 'Accounts',
                label: '128',
                icon: Wallet,
                variant: 'default',
              },
              {
                title: 'Real State',
                label: '9',
                icon: Building,
                variant: 'ghost',
              },
              {
                title: 'Stocks',
                label: '',
                icon: ActivityIcon,
                variant: 'ghost',
              },
              {
                title: 'Loan',
                label: '23',
                icon: ReceiptIndianRupee,
                variant: 'ghost',
              },
              {
                title: 'Trash',
                label: '',
                icon: Trash2,
                variant: 'ghost',
              },
              {
                title: 'Archive',
                label: '',
                icon: Archive,
                variant: 'ghost',
              },
            ]}
          />
        </ResizablePanel>
        <ResizableHandle withHandle />
        <ResizablePanel defaultSize={defaultLayout[1]} minSize={15}>
          <OverViewSwitch />
        </ResizablePanel>
        <ResizableHandle withHandle />

        <ResizablePanel defaultSize={defaultLayout[2]} minSize={15}>
          asasasas
        </ResizablePanel>
      </ResizablePanelGroup>
    </TooltipProvider>
  );
};

export default SideBar;
