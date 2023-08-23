/**
 * OPERA Cloud Price Availability Rate API
 * APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { ImageStyleType } from './ImageStyleType';
/**
 * Represents on image set record.
 * @export
 * @interface ImageSetType
 */
export interface ImageSetType {
    /**
     * Language identification.
     * @type {string}
     * @memberof ImageSetType
     */
    language?: string;
    /**
     * The image set name.
     * @type {string}
     * @memberof ImageSetType
     */
    imageSet?: string;
    /**
     * The image set sequence ID.
     * @type {number}
     * @memberof ImageSetType
     */
    sequenceId?: number;
    /**
     *
     * @type {ImageStyleType}
     * @memberof ImageSetType
     */
    imageStyle?: ImageStyleType;
    /**
     * The image set type.
     * @type {string}
     * @memberof ImageSetType
     */
    imageType?: string;
    /**
     * The image set hotel code.
     * @type {string}
     * @memberof ImageSetType
     */
    hotelId?: string;
    /**
     * The image set chain code.
     * @type {string}
     * @memberof ImageSetType
     */
    chainCode?: string;
    /**
     * The image set URL.
     * @type {string}
     * @memberof ImageSetType
     */
    imageURL?: string;
    /**
     * The image set description.
     * @type {string}
     * @memberof ImageSetType
     */
    description?: string;
    /**
     * The image set order.
     * @type {number}
     * @memberof ImageSetType
     */
    imageOrder?: number;
    /**
     * The image set website.
     * @type {string}
     * @memberof ImageSetType
     */
    website?: string;
}
/**
 * Check if a given object implements the ImageSetType interface.
 */
export declare function instanceOfImageSetType(value: object): boolean;
export declare function ImageSetTypeFromJSON(json: any): ImageSetType;
export declare function ImageSetTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ImageSetType;
export declare function ImageSetTypeToJSON(value?: ImageSetType | null): any;
