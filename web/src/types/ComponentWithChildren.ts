import { Component, JSX } from 'solid-js';

export type ComponentWithChildren<P = {}> = Component<
  { children: JSX.Element } & P
>;
