import Provider from "../provider/Provider";
import Entries from "./Entries";
import { PageHeader, Row, Col, Layout } from "antd";
import { Content } from "antd/lib/layout/layout";
import AddEntry from "./AddEntry";

const App = () => {
  return (
    <Layout style={{ height: "100%", overflow: "hidden", overflowY: "auto" }}>
      <Provider>
        <Row justify="center" gutter={[8, 32]}>
          <Col span={24}>
            <PageHeader title="Mini-system information" subTitle="WORK SMART" />
          </Col>
          <Col span={22}>
            <Content>
              <Row gutter={[8, 8]}>
                <Col span={24}>
                  <AddEntry />
                </Col>
                <Col span={24}>
                  <Entries />
                </Col>
              </Row>
            </Content>
          </Col>
        </Row>
      </Provider>
    </Layout>
  );
};

export default App;
