import cairosvg
import argparse
import os

def convert_svg_to_png(svg_path, png_path):
    cairosvg.svg2png(url=svg_path, write_to=png_path)

if __name__ == "__main__":
    # 创建命令行参数解析器
    parser = argparse.ArgumentParser(description='Convert SVG to PNG.')
    parser.add_argument('svg_path', type=str, help='Path to the input SVG file.')

    # 解析命令行参数
    args = parser.parse_args()

    # 提取SVG文件的路径
    svg_path = args.svg_path
    # 改变文件后缀为.png
    png_path = os.path.splitext(svg_path)[0] + '.png'
    
    # 转换SVG文件到PNG
    convert_svg_to_png(svg_path, png_path)
