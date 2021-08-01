import * as React from 'react';
import './DashboardLayout.css';

export function DashboardLayout({
  sidebar,
  children,
}: {
  sidebar?: React.ReactElement;
  children: React.ReactElement;
}): React.ReactElement {
  return (
    <section className="DashboardLayout_Root">
      <aside className="DashboardLayout_SidebarContainer">{sidebar}</aside>
      <main className="DashboardLayout_MainContainer">{children}</main>
    </section>
  );
}
