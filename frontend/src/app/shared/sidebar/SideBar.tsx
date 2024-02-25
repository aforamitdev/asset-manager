import { cn } from '@/lib/utils';
import { Nav } from '@/shadcn/nav';
import {
  ResizableHandle,
  ResizablePanel,
  ResizablePanelGroup,
} from '@/shadcn/resizable';
import { Separator } from '@/shadcn/separator';
import { TooltipProvider } from '@/shadcn/tooltip';
import { Archive, ArchiveX, File, Inbox, Send, Trash2 } from 'lucide-react';
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
              'flex h-[52px] items-center justify-center',
              isCollapsed ? 'h-[52px]' : 'px-2'
            )}
          >
            account
          </div>
          <Separator />
          <Nav
            isCollapsed={isCollapsed}
            links={[
              {
                title: 'Inbox',
                label: '128',
                icon: Inbox,
                variant: 'default',
              },
              {
                title: 'Drafts',
                label: '9',
                icon: File,
                variant: 'ghost',
              },
              {
                title: 'Sent',
                label: '',
                icon: Send,
                variant: 'ghost',
              },
              {
                title: 'Junk',
                label: '23',
                icon: ArchiveX,
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
          asas
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
