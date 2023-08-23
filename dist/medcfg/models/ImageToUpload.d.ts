/**
 * OPERA Cloud Content Service
 * Opera Content Service offers capability to manage large content such as images and files.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { ImageUploadType } from './ImageUploadType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Request to upload an image.
 * @export
 * @interface ImageToUpload
 */
export interface ImageToUpload {
    /**
     * List of Image details to upload including image to upload.
     * @type {Array<ImageUploadType>}
     * @memberof ImageToUpload
     */
    images?: Array<ImageUploadType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof ImageToUpload
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof ImageToUpload
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the ImageToUpload interface.
 */
export declare function instanceOfImageToUpload(value: object): boolean;
export declare function ImageToUploadFromJSON(json: any): ImageToUpload;
export declare function ImageToUploadFromJSONTyped(json: any, ignoreDiscriminator: boolean): ImageToUpload;
export declare function ImageToUploadToJSON(value?: ImageToUpload | null): any;
