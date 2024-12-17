import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Auth from "./pages/auth";
import Base from "./pages/base";
import CreateQuiz from "./pages/create_quiz";
import QrQuiz from "./pages/qr_quiz";
import UserPanel from "./pages/user_panel";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/">
          <Route path="authorization" element={<Auth />} />
          <Route path="" element={<Base />} />
          <Route path="create_quiz" element={<CreateQuiz />} />
          <Route path="qr" element={<QrQuiz />} />
          <Route path="panel" element={<UserPanel />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
