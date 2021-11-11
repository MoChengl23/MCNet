namespace KCPNET
{
    public class Msg
    {
        public string name;
        public CMD cmd;
       
    }

    public enum CMD{
        Ping,
        ChatMessage,
    }
    
}