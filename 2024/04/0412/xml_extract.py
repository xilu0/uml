import xml.etree.ElementTree as ET

def extract_objs_content(clixml_data) -> str:
    # 查找<Objs标签开始的位置
    start_index = clixml_data.find('<Objs')
    if start_index == -1:
        return "No <Objs> tag found."

    # 查找</Objs>标签结束的位置
    end_index = clixml_data.find('</Objs>', start_index)
    if end_index == -1:
        return "No </Objs> tag found."

    # 计算</Objs>标签闭合部分的位置，加上7是因为"</Objs>"的长度
    end_index += len('</Objs>')

    # 返回从<Objs>到</Objs>之间的内容
    return clixml_data[start_index:end_index]

def parse_clixml(clixml_data):
    # 创建命名空间字典，因为CLIXML使用了命名空间
    namespaces = {'ps': 'http://schemas.microsoft.com/powershell/2004/04'}

    # 解析 XML
    root = ET.fromstring(clixml_data)

    results = {
        'Action Messages': [],
        'Statuses': [],
        'Error Messages': []
    }

    # 遍历所有的Obj标签，处理进度信息
    for obj in root.findall('ps:Obj', namespaces):
        for ms in obj.findall('ps:MS', namespaces):
            action_msg = ms.find('.//ps:AV', namespaces)
            status = ms.find('.//ps:T', namespaces)
            if action_msg is not None:
                results['Action Messages'].append(action_msg.text)
            if status is not None:
                results['Statuses'].append(status.text)

    # 遍历所有错误信息
    for error in root.findall('ps:S', namespaces):
        if error.attrib['S'] == 'Error':
            results['Error Messages'].append(error.text)

    return results

# 示例使用
clixml_data = '''
CLIXML  
<Objs Version="1.1.0.1" xmlns="http://schemas.microsoft.com/powershell/2004/04">
    <Obj S="progress" RefId="0">
        <TN RefId="0">
            <T>System.Management.Automation.PSCustomObject</T>
            <T>System.Object</T>
        </TN>
        <MS>
            <I64 N="SourceId">1</I64>
            <PR N="Record">
                <AV>Preparing modules for first use.</AV>
                <AI>0</AI>
                <Nil/>
                <PI>-1</PI>
                <PC>-1</PC>
                <T>Completed</T>
                <SR>-1</SR>
                <SD> </SD>
            </PR>
        </MS>
    </Obj>
    <S S="Error">Set-ADAccountPassword : The specified network password is not correct</S>
</Objs>
'''

results = parse_clixml(extract_objs_content(clixml_data))
print(results)
