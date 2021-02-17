import React from "react";
import { Form, Header } from "semantic-ui-react";
import logo from "../../logo.svg";

interface Props {
    header: string;
    children: Array<JSX.Element>;
}

export const AuthLayout: React.FC<Props> = (props: Props) => (
    <div className="auth-main">
        <div className="auth-content">
            <div className="auth-card">
                <img src={logo} alt="Logo" className="auth-logo" />
                <Header as="h2" color="black" textAlign="center">
                    {props.header}
                </Header>
                <Form.Group
                    size="large"
                    className="auth-form"
                    autoComplete="off"
                >
                    {props.children}
                </Form.Group>
            </div>
        </div>
    </div>
);
