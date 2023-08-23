/**
 * OPERA Cloud Activity Management API
 * APIs to cater for Activity Configuration functionality in OPERA Cloud. In this module you can retrieve, create, update Activity configuration codes, for example create a new Activity Type.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { TranslationTextType2000 } from './TranslationTextType2000';
/**
 * Contains Common Master configuration detail.
 * @export
 * @interface ActivityResultsConfigDetailType
 */
export interface ActivityResultsConfigDetailType {
    /**
     * Common Master unique code.
     * @type {string}
     * @memberof ActivityResultsConfigDetailType
     */
    code?: string;
    /**
     *
     * @type {TranslationTextType2000}
     * @memberof ActivityResultsConfigDetailType
     */
    description?: TranslationTextType2000;
    /**
     * Common Master record sequence number.
     * @type {number}
     * @memberof ActivityResultsConfigDetailType
     */
    displayOrder?: number;
    /**
     * Indicates the Activity Results is inactive or not.
     * @type {boolean}
     * @memberof ActivityResultsConfigDetailType
     */
    inactive?: boolean;
}
/**
 * Check if a given object implements the ActivityResultsConfigDetailType interface.
 */
export declare function instanceOfActivityResultsConfigDetailType(value: object): boolean;
export declare function ActivityResultsConfigDetailTypeFromJSON(json: any): ActivityResultsConfigDetailType;
export declare function ActivityResultsConfigDetailTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ActivityResultsConfigDetailType;
export declare function ActivityResultsConfigDetailTypeToJSON(value?: ActivityResultsConfigDetailType | null): any;
