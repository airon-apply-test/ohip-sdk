/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Sales Event Management API
 * APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


/**
 * An enumeration of image styles.
 * @export
 */
export const ImageStyleType = {
    Full: 'Full',
    Thumbnail: 'Thumbnail',
    Icon: 'Icon',
    Square: 'Square',
    Banner: 'Banner',
    Button: 'Button',
    Lowres: 'Lowres',
    Highres: 'Highres'
} as const;
export type ImageStyleType = typeof ImageStyleType[keyof typeof ImageStyleType];


export function ImageStyleTypeFromJSON(json: any): ImageStyleType {
    return ImageStyleTypeFromJSONTyped(json, false);
}

export function ImageStyleTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ImageStyleType {
    return json as ImageStyleType;
}

export function ImageStyleTypeToJSON(value?: ImageStyleType | null): any {
    return value as any;
}

