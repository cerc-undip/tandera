import React from "react";
import {Button, Container, Header, Menu, Segment} from "semantic-ui-react";

export const LandingPage: React.FC = () => (
    <Segment
        inverted
        textAlign='center'
        style={{ minHeight: 700, padding: '1em 0em' }}
        vertical
    >
        <Menu
            //fixed={fixed ? 'top' : null}
            inverted={true}
            pointing={true}
            secondary={true}
            size='large'
        >
            <Container>
                <Menu.Item as='a' active>
                    Home
                </Menu.Item>
                <Menu.Item as='a'>Work</Menu.Item>
                <Menu.Item as='a'>Company</Menu.Item>
                <Menu.Item as='a'>Careers</Menu.Item>
                <Menu.Item position='right'>
                    <Button as='a' inverted>
                        Log in
                    </Button>
                </Menu.Item>
            </Container>
        </Menu>
        <Container text>
            <Header
                as='h1'
                content='Tandera'
                inverted
                style={{
                    fontSize: '4em',
                    fontWeight: 'normal',
                    marginBottom: 0,
                    marginTop: '3em',
                }}
            />
            <Header
                as='h2'
                content='Tangkap bendera! Haha get it?'
                inverted
                style={{
                    fontSize: '1.7em',
                    fontWeight: 'normal',
                    marginTop: '1.5em',
                }}
            />
        </Container>
    </Segment>
);