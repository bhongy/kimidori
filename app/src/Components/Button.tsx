import * as React from 'react';
import clsx from 'clsx';
import './Button.css';

// BaseButton resets user-agent styles and provide "bare" defaults
export function BaseButton({
  children,
  className,
}: {
  children: React.ReactNode;
  className?: string;
}): React.ReactElement {
  return <button className={clsx('BaseButton', className)}>{children}</button>;
}

// BlockButton fills its container's width
export function BlockButton({
  className,
  children,
}: {
  children: React.ReactNode;
  className?: string;
}): React.ReactElement {
  return (
    <BaseButton className={clsx('BlockButton', className)}>
      {children}
    </BaseButton>
  );
}
