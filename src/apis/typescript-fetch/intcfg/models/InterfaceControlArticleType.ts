/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Integration Configuration API
 * APIs catering to Integration Configuration in OPERA Cloud.  Operations such as get Hotel Interface Types, or get UDF mappings can be found in this module.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface InterfaceControlArticleType
 */
export interface InterfaceControlArticleType {
    /**
     * Hotel Code of the hotel interface.
     * @type {string}
     * @memberof InterfaceControlArticleType
     */
    hotelId?: string;
    /**
     * Logo of the hotel interface.
     * @type {string}
     * @memberof InterfaceControlArticleType
     */
    logo?: string;
    /**
     * Internal code of a translation article setup.
     * @type {number}
     * @memberof InterfaceControlArticleType
     */
    internalCode?: number;
    /**
     * Selector of a translation article setup.
     * @type {string}
     * @memberof InterfaceControlArticleType
     */
    selector?: string;
    /**
     * Article Number of a translation article setup.
     * @type {string}
     * @memberof InterfaceControlArticleType
     */
    articleNo?: string;
    /**
     * Description of a translation article setup.
     * @type {string}
     * @memberof InterfaceControlArticleType
     */
    description?: string;
}

/**
 * Check if a given object implements the InterfaceControlArticleType interface.
 */
export function instanceOfInterfaceControlArticleType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function InterfaceControlArticleTypeFromJSON(json: any): InterfaceControlArticleType {
    return InterfaceControlArticleTypeFromJSONTyped(json, false);
}

export function InterfaceControlArticleTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): InterfaceControlArticleType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'logo': !exists(json, 'logo') ? undefined : json['logo'],
        'internalCode': !exists(json, 'internalCode') ? undefined : json['internalCode'],
        'selector': !exists(json, 'selector') ? undefined : json['selector'],
        'articleNo': !exists(json, 'articleNo') ? undefined : json['articleNo'],
        'description': !exists(json, 'description') ? undefined : json['description'],
    };
}

export function InterfaceControlArticleTypeToJSON(value?: InterfaceControlArticleType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'hotelId': value.hotelId,
        'logo': value.logo,
        'internalCode': value.internalCode,
        'selector': value.selector,
        'articleNo': value.articleNo,
        'description': value.description,
    };
}

