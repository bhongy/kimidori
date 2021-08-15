import * as React from 'react';
import clsx from 'clsx';
import './Button.css';

type BaseButtonProps = {
  children: React.ReactNode;
  className?: string;
  onClick?: React.MouseEventHandler<HTMLButtonElement>;
};

// BaseButton resets user-agent styles and provide "bare" defaults
export function BaseButton({
  children,
  className,
  onClick,
}: BaseButtonProps): React.ReactElement {
  return (
    <button
      className={clsx('BaseButton', className)}
      onMouseUp={({currentTarget}) => currentTarget.blur()}
      onClick={onClick}
    >
      {children}
    </button>
  );
}

type BlockButtonProps = BaseButtonProps;

// BlockButton fills its container's width
export function BlockButton({
  className,
  ...passThroughProps
}: BlockButtonProps): React.ReactElement {
  return (
    <BaseButton
      {...passThroughProps}
      className={clsx('BlockButton', className)}
    />
  );
}
